package main

import (
	"sync"
)

type PixelPipeFan []PixelPipe

func (pip PixelPipe) Collect(done <-chan bool, num int) (result []Pixel) {
	for i := 0; i < num; i++ {
		select {
		case val := <-pip:
			result = append(result, val)
		case <-done:
			return
		}
	}
	return
}
func (pip PixelPipe) Fan(done <-chan bool, workers int, fn func(Pixel) Pixel) PixelPipe {
	return pip.FanOut(done, workers, fn).FanIn(done)
}
func (pip PixelPipe) FanOut(done <-chan bool, workers int, fn func(Pixel) Pixel) PixelPipeFan {
	fan := PixelPipeFan{}
	for i := 0; i < workers; i++ {
		fan = append(fan, pip.worker(done, fn))
	}
	return fan
}
func (pip PixelPipe) Filter(done <-chan bool, fn func(Pixel) bool) PixelPipe {
	out := make(chan Pixel)
	go func() {
		defer close(out)
		for val := range pip {
			if fn(val) {
				select {
				case out <- val:
				case <-done:
					return
				}
			}
		}
	}()
	return out
}
func (pip PixelPipe) Pipe(done <-chan bool, fn func(Pixel) Pixel) PixelPipe {
	out := make(chan Pixel)
	go func() {
		defer close(out)
		for val := range pip {
			select {
			case out <- fn(val):
			case <-done:
				return
			}
		}
	}()
	return out
}
func (pip PixelPipe) worker(done <-chan bool, fn func(Pixel) Pixel) PixelPipe {
	out := make(chan Pixel)
	go func() {
		defer close(out)
		for val := range pip {
			select {
			case out <- fn(val):
			case <-done:
				return
			}
		}
	}()
	return out
}
func (fan PixelPipeFan) FanIn(done <-chan bool) PixelPipe {
	var wg sync.WaitGroup
	out := make(chan Pixel)
	output := func(pl PixelPipe) {
		defer wg.Done()
		for val := range pl {
			select {
			case out <- val:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(fan))
	for _, val := range fan {
		go output(val)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func (fan PixelPipeFan) Filter(done <-chan bool, fn func(Pixel) bool) (result PixelPipeFan) {
	for _, pipe := range fan {
		result = append(result, pipe.Filter(done, fn))
	}
	return
}
func (fan PixelPipeFan) Pipe(done <-chan bool, fn func(Pixel) Pixel) (result PixelPipeFan) {
	for _, pipe := range fan {
		result = append(result, pipe.Pipe(done, fn))
	}
	return
}
