package main

import (
    "flag"
    "os"

    "k8s.io/kubernetes/pkg/scheduler/framework/runtime"
    "k8s.io/kubernetes/pkg/scheduler/framework/plugins/legacyregistry"

    "custom_scheduler/pkg/scheduler"
)

func main() {
    framework.RegisterPlugin(scheduler.Name, scheduler.New)

    // Initialize the scheduler framework
    command := app.NewSchedulerCommand(
        app.WithPlugin(scheduler.Name, scheduler.New),
    )

    command.Flags().AddGoFlagSet(flag.CommandLine)
    if err := command.Execute(); err != nil {
        os.Exit(1)
    }
}
