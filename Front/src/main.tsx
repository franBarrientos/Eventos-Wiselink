import ReactDOM from 'react-dom/client'
import {RouterProvider} from "react-router-dom";
import router from "./router.tsx";
import './index.css'
import {ToastContainer} from "react-toastify";
import 'react-toastify/dist/ReactToastify.css';
import {AppContextProvider} from "./context/AppProvider.tsx";


ReactDOM.createRoot(document.getElementById('root')!).render(
    <>
        <AppContextProvider>
            <RouterProvider router={router}/>
        </AppContextProvider>
        <ToastContainer/>

    </>
)
