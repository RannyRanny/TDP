import React, { useEffect, useState } from 'react';
import { fetchTasksForUser } from '../api/taskApi';

interface Task {
    id: number;
    title: string;
    description: string;
}

interface TaskListProps {
    userId: number;
}

const TaskList: React.FC<TaskListProps> = ({ userId }) => {
    const [tasks, setTasks] = useState<Task[]>([]);

    useEffect(() => {
        const fetchTasks = async () => {
            const tasksData = await fetchTasksForUser(userId);
            setTasks(tasksData);
        };

        fetchTasks();
    }, [userId]);

    return (
        <ul>
            {tasks.map((task) => (
                    <li key={task.id}>{task.title} - {task.description}</li>
                ))}
        </ul>
    );
}

export default TaskList;
