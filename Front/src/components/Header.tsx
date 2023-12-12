import { FunctionComponent } from "react"
import { ArrowRight } from "./../icons/ArrowRight"
import useApp from "../hooks/useApp.tsx"
import { useNavigate } from "react-router-dom"
import { toast } from "react-toastify"

interface OwnProps {}

type Props = OwnProps

const Header: FunctionComponent<Props> = () => {
  const { user, cleanUser } = useApp()
  const navigate = useNavigate()
  const logout = () => {
    toast.success("Logout success")
    cleanUser()
    navigate("/")
  }

  return (
    <header className="text-gray-600 body-font shadow-2xl">
      <div className="container mx-auto flex flex-wrap p-5 flex-col md:flex-row items-center">
        <a className="flex title-font font-medium items-center text-gray-900 mb-4 md:mb-0">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            className="w-10 h-10 text-white p-2 bg-indigo-500 rounded-full"
            viewBox="0 0 24 24"
          >
            <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"></path>
          </svg>
          <span className="ml-3 text-xl">Wiselink Events</span>
        </a>
        <nav className="md:ml-auto flex flex-wrap items-center text-base justify-center">
          <a onClick={() => navigate(user && user.Role == "ADMIN" ? "/admin" : "/")} className=" bg-gray-100 border-0 rounded py-1 px-3 focus:outline-none hover:bg-gray-200  mr-5 hover:text-gray-900 cursor-pointer">Home</a>
          {user && user.Role == "USER" && <a onClick={() => navigate("/events")} className=" bg-gray-100 border-0 rounded py-1 px-3 focus:outline-none hover:bg-gray-200  mr-5 hover:text-gray-900 cursor-pointer">My Events</a>}
        </nav>
        {user ? (
          <button
            onClick={logout}
            className="inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0"
          >
            Logout
            <ArrowRight />
          </button>
        ) : (
          <button
            onClick={() => navigate("/login")}
            className="inline-flex items-center bg-gray-100 border-0 py-1 px-3 focus:outline-none hover:bg-gray-200 rounded text-base mt-4 md:mt-0"
          >
            Login
            <ArrowRight />
          </button>
        )}
      </div>
    </header>
  )
}

export default Header
