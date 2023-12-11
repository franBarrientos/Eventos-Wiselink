import {createBrowserRouter} from "react-router-dom";
import LayoutHome from "./layouts/LayoutHome.tsx";
import Home from "./pages/Home.tsx";
import LayoutAuth from "./layouts/LayoutAuth.tsx";
import Login from "./pages/Login.tsx";
import Register from "./pages/Register.tsx";
import MyEvents from "./pages/MyEvents.tsx";
import Admin from "./pages/Admin.tsx";

const router = createBrowserRouter([
    {
        path: "/",
        element: <LayoutHome />,
        children: [
            {
                index: true,
                element: <Home />,
            },
        ],
    },
    {
        path: "/events",
        element: <LayoutHome />,
        children: [
            {
                index: true,
                element: <MyEvents />,
            },
        ],
    },
    {
        path: "/admin",
        element: <LayoutHome />,
        children: [
            {
                index: true,
                element: <Admin />,
            },
        ],
    },
    {
        path: "/login",
        element: <LayoutAuth />,
        children: [
            {
                index: true,
                element: <Login />,
            },
        ],
    },
    {
        path: "/register",
        element: <LayoutAuth />,
        children: [
            {
                index: true,
                element: <Register />,
            },
        ],
    }
]);

export default router;