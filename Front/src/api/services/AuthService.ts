import {LoginDTO, UserAddDTO} from "../dtos/input.ts";
import {clienteAxios} from "../client.ts";


export const LoginRequest = (data: LoginDTO) => {
    return clienteAxios.post("/auth/login", data, {
        headers: {
            'Content-Type': 'application/json',
        },
    })
}


export const RegisterRequest = (data: UserAddDTO) => {
    return clienteAxios.post("/auth/register", data, {
        headers: {
            'Content-Type': 'application/json',
        },
    })
}