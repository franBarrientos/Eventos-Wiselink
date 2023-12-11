import  {FunctionComponent} from 'react';
import Header from "../components/Header.tsx";
import {Outlet} from "react-router-dom";
import Footer from "../components/Footer.tsx";

interface OwnProps {
}

type Props = OwnProps;

const LayoutHome: FunctionComponent<Props> = () => {

    return (
        <div className=" bg-gray-100 flex min-h-screen flex-col justify-between">
            <Header/>
            <Outlet/>
            <Footer />
        </div>

    );
};

export default LayoutHome;
