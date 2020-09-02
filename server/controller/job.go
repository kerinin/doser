package controller

import (
	"context"
	"sync"
)

type Job struct {
	cancelFunc context.CancelFunc
	wg         *sync.WaitGroup
}
