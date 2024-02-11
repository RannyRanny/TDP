import React from 'react';
import TaskList from "../../components/TaskList";
import {ViewBase} from "../../shared/components/view-base/view-base";

const Main = () => {
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
            список дел
        </ViewBase>
    );
}

export default Main
