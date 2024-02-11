import { createBrowserRouter } from "react-router-dom";
import Main from "./views/main";

export const routes = createBrowserRouter([{
    path: '/',
    element: <Main />
}])
