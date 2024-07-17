package scheduler

import (
    "k8s.io/apimachinery/pkg/runtime"
    "k8s.io/kubernetes/pkg/scheduler/framework"
)

type Config struct {
    ScoreWeight int `json:"scoreWeight,omitempty"`
}

func NewConfig() runtime.Object {
    return &Config{}
}
