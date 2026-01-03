Cancellabe Http Job Server

- The Handler receives intent
    {
        num1: 10
        num2: 20
        timeout: "5s"
    }

- The Service layer descides policy (Orchestration Layer)
    ctx, cancel := context.WithTimeout(parentCtx, input.Timeout)
    defer cancel()

- The Execution Layer owns the mechanics in which the job is carried out (Domain - Job Logic)
    const stepDuration = 100 * time.Millisecond
