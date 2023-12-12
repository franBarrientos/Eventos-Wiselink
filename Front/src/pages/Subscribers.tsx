import  { FunctionComponent, useEffect, useState } from "react"
import CardSkeleton from "../components/CardSkeleton.tsx"
import { useParams } from "react-router-dom"
import { GetSubscribersToEvent } from "../api/services/EventsService.ts"
import { toast } from "react-toastify"
import UserCard from "../components/UserCard.tsx";
import {UserDTO} from "../api/dtos/output.ts";
import Pagination from "../components/Pagination.tsx";
import {pageToSkipSubject} from "../utils/subjectsRx.ts";

interface OwnProps {}

type Props = OwnProps

const Subscribers: FunctionComponent<Props> = () => {
  const [isLoading, setIsLoading] = useState(true)
  const { id } = useParams()
  const [subscribers, setSubscribers] = useState<UserDTO[]>([])
  const [currentPage, setCurrentPage] = useState(1)

  const GetSubscribers = (id: number, page: number = 1) => {
    GetSubscribersToEvent(id, page)
        .then((r) => {
          setSubscribers(r)
          setTimeout(() => {
            setIsLoading(false)
          }, 1000)
        })
        .catch(() => {
          toast.error("Something went wrong")
        })
  }


  useEffect(() => {
    pageToSkipSubject.getSubject.subscribe( (value) => {
      setCurrentPage(value)
      GetSubscribers(Number(id), value)
    })
  }, []);


  useEffect(() => {
    GetSubscribers(Number(id) )
  }, [])

  return (
    <section className="text-gray-600 body-font">
      <div className="container px-5 py-8 mx-auto">
        <div className="flex flex-wrap -mx-4  -my-8">
          {isLoading && (
            <>
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
            </>
          )}
        </div>
        <section className="text-gray-600 body-font">
          <div className="container px-5 py-24 mx-auto">
            <div className="flex flex-col text-center w-full mb-8">
              <h1 className="sm:text-3xl text-2xl font-medium title-font mb-4 text-gray-900">
                Subscribers
              </h1>
            </div>
            <div className="flex flex-wrap -m-2">
              <div className="p-2 lg:w-1/3 md:w-1/2 w-full">

                {subscribers?.map((s) => (
                    <UserCard key={s.Id} {...s} />
                ))}

              </div>
            </div>
          </div>
        </section>
         <div className="mx-auto w-full flex justify-center mt-10">
              <Pagination currentPage={currentPage} />
          </div>
      </div>
    </section>
  )
}

export default Subscribers
