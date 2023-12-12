import  { FunctionComponent } from "react"
import {UserDTO} from "../api/dtos/output.ts";


type Props = UserDTO

const UserCard: FunctionComponent<Props> = (props) => {
  return (
    <div className="h-full flex items-center border-gray-200 border p-1 rounded-lg">
      <img
        alt="blog"
        src="https://static.vecteezy.com/system/resources/thumbnails/009/734/564/small/default-avatar-profile-icon-of-social-media-user-vector.jpg"
        className=" mr-4 w-8 h-4 rounded-full flex-shrink-0 object-cover object-center"
      />
      <div className="flex-grow">
        <h2 className="text-gray-900 title-font font-medium">
          {props.FirstName} {props.LastName}
        </h2>
        <p className="text-gray-500">{props.Email}</p>
      </div>
    </div>
  )
}

export default UserCard
