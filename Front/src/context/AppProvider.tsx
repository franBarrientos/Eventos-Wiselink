import {UserDTO} from "../api/dtos/output.ts";
import {createContext, ReactNode, useEffect, useState} from "react";

interface MyContextType {
    user: UserDTO;
    setUser: (user: UserDTO) => void;
    cleanUser: () => void;
}


export const AppContext = createContext<MyContextType>({} as MyContextType);

type AppContextProviderProps = {
    children: ReactNode;
};

export const AppContextProvider: React.FC<AppContextProviderProps> = ({children}) => {

    const [user, setUser] = useState<UserDTO>(null!);

    const setUserContext = (user: UserDTO) => {
        setUser(user);
        localStorage.setItem("user", JSON.stringify(user));
    }

    useEffect(() => {
        const user = localStorage.getItem("user");
        if (user) {
            setUser(JSON.parse(user));
        }
    }, []);

    const cleanUser = () => {
        setUser(null!);
        localStorage.removeItem("user");
        localStorage.removeItem("token")

    }

    const contextValue: MyContextType = {
        user,
        setUser: setUserContext,
        cleanUser

    };

    return (
        <AppContext.Provider value={contextValue}>{children}</AppContext.Provider>
    );
};

export default AppContext;