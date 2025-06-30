package models

import (
	"sync"
)

type Storage struct {
	Db map[string]*Task
	Mu sync.RWMutex
}
