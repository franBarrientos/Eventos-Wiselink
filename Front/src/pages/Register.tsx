import {FunctionComponent} from 'react';
import {useNavigate} from "react-router-dom";
import { UserAddDTO} from "../api/dtos/input.ts";
import { RegisterRequest} from "../api/services/AuthService.ts";
import {toast} from "react-toastify";
import useApp from "../hooks/useApp.tsx";
import {FieldValues, SubmitHandler, useForm} from "react-hook-form";

interface OwnProps {
}

type Props = OwnProps;

const Register: FunctionComponent<Props> = () => {

    const navigate = useNavigate();
    const {setUser} = useApp()
    const {register, handleSubmit} = useForm();

    const handlerRegister = (data: UserAddDTO) => {
        RegisterRequest(data).then(r => {
            toast.success("Register success");
            const redirect = r.data.User.Role === "ADMIN" ? "/admin" : "/";
            setUser(r.data.User);
            localStorage.setItem("token", r.data.Token.AccessToken)
            navigate(redirect);
        }).catch((e) => {
            toast.error("Register failed");
            console.log(e)
        })
    }

    return (
        <div className="bg-white shadow w-full rounded-lg divide-y divide-gray-200">

            <div className="px-5 py-7">
                <h1 className="font-semibold text-center text-2xl mb-5">Register</h1>
                <form onSubmit={handleSubmit(handlerRegister as SubmitHandler<FieldValues>)}>
                <label className="font-semibold text-md text-gray-600 pb-1 block">Firstname</label>
                <input {...register('FirstName', {required: true})}  type="text" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"/>
                <label className="font-semibold text-md text-gray-600 pb-1 block">Lastname</label>
                <input {...register('LastName', {required: true})}  type="text" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"/>
                <label className="font-semibold text-md text-gray-600 pb-1 block">E-mail</label>
                <input {...register('Email', {required: true})} type="text" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"/>
                <label className="font-semibold text-md text-gray-600 pb-1 block">Password</label>
                <input {...register('Password', {required: true})} type="password" className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"/>
                <button type="submit"
                        className="transition duration-200 bg-indigo-500 hover:bg-indigo-600 focus:bg-indigo-700 focus:shadow-sm focus:ring-4 focus:ring-blue-500 focus:ring-opacity-50 text-white w-full py-2.5 rounded-lg text-sm shadow-sm hover:shadow-md font-semibold text-center inline-block">
                    <span className="inline-block mr-2">Register</span>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor"
                         className="w-4 h-4 inline-block">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                              d="M17 8l4 4m0 0l-4 4m4-4H3"/>
                    </svg>
                </button>
                </form>
            </div>
            <div className="py-5">
                <div className="flex justify-center gap-1">
                    <div className="text-center sm:text-left whitespace-nowrap">
                        <button
                            onClick={() => navigate("/login")}
                            className="transition duration-200 mx-5 px-5 py-4 cursor-pointer font-normal text-sm rounded-lg text-gray-500 hover:bg-gray-100 focus:outline-none focus:bg-gray-200 focus:ring-2 focus:ring-gray-400 focus:ring-opacity-50 ring-inset">
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                                 stroke="currentColor" className="w-4 h-4 inline-block align-text-top">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                      d="M8 11V7a4 4 0 118 0m-4 8v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2z"/>
                            </svg>
                            <span className="inline-block ml-3 ">You already have an account? <span
                                className="ml-2  font-semibold text-indigo-500"> Login</span></span>
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Register;
