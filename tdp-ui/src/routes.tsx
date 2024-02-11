import { createBrowserRouter } from 'react-router-dom';
import DailyDashboard from './views/daily-dashboard';
import Settings from './views/settings';

export const routes = createBrowserRouter([{
    path: '/',
    element: <DailyDashboard />,
}, {
    path: '/settings',
    element: <Settings />,
}])
