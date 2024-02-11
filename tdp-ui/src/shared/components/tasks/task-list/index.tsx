import React from "react";
import { TaskListItemComponent } from "../task-list-item";

const mockTasks = [{
    taskId: '1',
    title: 'покекать',
    taskHour: '07:00',
    done: false,
}]

export const TaskList = () => {
    return (
        <div className='task-list'>
            { mockTasks.map((task) => (
                <TaskListItemComponent
                    key={ task.taskId }
                    taskId={ task.taskId }
                    taskHour={ task.taskHour }
                    done={ task.done }
                    title={ task.title }
                />
            ))}
        </div>
    )
}
