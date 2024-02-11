import React from 'react';
import { ViewBase } from '../../shared/components/view-base/view-base';
import { TaskList } from '../../shared/components/tasks/task-list';

const DailyDashboard = () => {
    const today = new Date().toLocaleDateString('ru-Ru', { weekday: 'short', month: 'short', day: 'numeric' });

    return (
        <ViewBase
            title={ today }
            showGoal={ true }
            footerOptions={{
                buttonText: 'Завершить день',
                buttonHandler: () => alert('пососи хуйца лошара'),
                showButton: true
            }}
        >
            <TaskList />
        </ViewBase>
    );
}

export default DailyDashboard
