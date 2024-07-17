package scheduler

import (
    "context"
    "fmt"

    "k8s.io/kubernetes/pkg/scheduler/framework"
    "k8s.io/kubernetes/pkg/scheduler/framework/plugins/names"
)

type DSLAppScheduler struct {
    handle framework.Handle
}

var _ framework.ScorePlugin = &DSLAppScheduler{}

const Name = "DSLAppScheduler"

// Name
func (d *DSLAppScheduler) Name() string {
    return Name
}

// Score 
func (d *DSLAppScheduler) Score(ctx context.Context, 
								state *framework.CycleState, 
								pod *v1.Pod, 
								nodeName string) (int64, *framework.Status) {
    nodeInfo, err := d.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
    if err != nil {
        return 0, framework.NewStatus(framework.Error, fmt.Sprintf("Failed to get node info: %v", err))
    }

    // Calculate score based on CPU, memory, and GPU availability
    score := calculateScore(nodeInfo)
    return score, framework.NewStatus(framework.Success)
}

// Score extension
func (d *DSLAppScheduler) ScoreExtensions() framework.ScoreExtensions {
    return d
}

// Normalize the scores if needed
func (d *DSLAppScheduler) NormalizeScore(ctx context.Context,
										 state *framework.CycleState,
										 pod *v1.Pod, 
										 scores framework.NodeScoreList) *framework.Status {
    return nil
}

// Returns the score
func calculateScore(nodeInfo *framework.NodeInfo) int64 {
    cpuScore := nodeInfo.Allocatable.MilliCPU
    memoryScore := nodeInfo.Allocatable.Memory
    gpuScore := nodeInfo.Allocatable.Extensions[ResourceGPU]
    return cpuScore + memoryScore + gpuScore
}

// create new instace
func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
    return &DSLAppScheduler{
        handle: h,
    }, nil
}

// Register
func init() {
    framework.RegisterPlugin(Name, New)
}
