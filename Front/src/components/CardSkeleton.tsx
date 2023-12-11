import { FunctionComponent } from "react"
import Skeleton, {SkeletonTheme} from "react-loading-skeleton"
import 'react-loading-skeleton/dist/skeleton.css'

interface OwnProps {}

type Props = OwnProps

const CardSkeleton: FunctionComponent<Props> = () => {
  return (
    <SkeletonTheme baseColor="#F3F3F3" highlightColor="#6366F1">
      <div className="py-8 px-3 lg:w-1/3">
      <div className={"flex py-8 px-3  h-full  rounded-2xl shadow-2xl "}>
        <div className="w-1/4 ml-2">
          <Skeleton borderRadius="10px" height={40} width={30}  />
        </div>
        <div className="w-3/4 flex flex-col gap-2">
          <Skeleton />
          <Skeleton />
          <Skeleton />
        </div>
      </div>
      </div>
    </SkeletonTheme>
  )
}

export default CardSkeleton
