
import './App.css'
import React from "react";
import TaskList from "./components/TaskList.tsx";

const App: React.FC = () => {
    const today = new Date().toLocaleDateString('en-US', { weekday: 'long', month: 'long', day: 'numeric' });

    return (
        <div className="App">
            <header>
                <h1>Tasks for {today}</h1>
            </header>
            <TaskList userId={1} /> {/* Example user ID */}
        </div>
    );
}

export default App
