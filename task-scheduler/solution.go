package taskscheduler

import "sort"

// time: O(m * n) where m = len(tasks) and n = n
// space: O(k) where k = type of tasks
func leastInterval(tasks []byte, n int) int {
	taskFreq := make(map[byte]int)
	for _, task := range tasks {
		taskFreq[task]++
	}

	interval := 0
	for len(taskFreq) > 0 {
		// List tasks to be completed
		tasksToComplete := []byte{}
		for k := range taskFreq {
			tasksToComplete = append(tasksToComplete, k)
		}

		// Prioritize tasks - select high frequency tasks first
		sort.Slice(tasksToComplete, func(i, j int) bool {
			return taskFreq[tasksToComplete[i]] > taskFreq[tasksToComplete[j]]
		})

		for i := 0; i <= n; i++ {
			if len(tasksToComplete) == 0 {
				if len(taskFreq) == 0 {
					break
				}
				interval++
				continue
			}

			task := tasksToComplete[0]
			tasksToComplete = tasksToComplete[1:]
			taskFreq[task]--
			if taskFreq[task] == 0 {
				delete(taskFreq, task)
			}

			interval++
		}

	}

	return interval
}
