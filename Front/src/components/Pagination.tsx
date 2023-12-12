import { FunctionComponent } from "react"
import { pageToSkipSubject } from "../utils/subjectsRx.ts"

interface OwnProps {
  currentPage: number
}

type Props = OwnProps

const Pagination: FunctionComponent<Props> = (props) => {
  return (
    <>
      <nav aria-label="Page navigation example">
        <div className="inline-flex -space-x-px text-sm">
          <button
            onClick={() => pageToSkipSubject.setSubject(props.currentPage - 1)}
            disabled={props.currentPage == 1}
            className="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-900 bg-white
                        border border-e-0 border-gray-300 rounded-s-lg hover:bg-gray-100 hover:text-gray-700"
          >
            Previous
          </button>
          <button
            onClick={() => pageToSkipSubject.setSubject(1)}
            className={`flex ${
              props.currentPage == 1
                ? `bg-indigo-400 text-white`
                : `bg-white text-gray-900`
            }  items-center justify-center px-3 h-8 leading-tight   border border-gray-300 hover:bg-gray-100 hover:text-gray-700`}
          >
            1
          </button>
          <button
            onClick={() => pageToSkipSubject.setSubject(2)}
            className={`flex ${
              props.currentPage == 2
                ? `bg-indigo-400 text-white`
                : `bg-white text-gray-900`
            }  items-center justify-center px-3 h-8 leading-tight   border border-gray-300 hover:bg-gray-100 hover:text-gray-700`}
          >
            2
          </button>
          <button
            onClick={() => pageToSkipSubject.setSubject(3)}
            className={`flex ${
              props.currentPage == 3
                ? `bg-indigo-400 text-white`
                : `bg-white text-gray-900`
            }  items-center justify-center px-3 h-8 leading-tight   border border-gray-300 hover:bg-gray-100 hover:text-gray-700`}
          >
            3
          </button>
          <button
            onClick={() => pageToSkipSubject.setSubject(4)}
            className={`flex ${
              props.currentPage == 4
                ? `bg-indigo-400 text-white`
                : `bg-white text-gray-900`
            }  items-center justify-center px-3 h-8 leading-tight   border border-gray-300 hover:bg-gray-100 hover:text-gray-700`}
          >
            4
          </button>
          <button
            onClick={() => pageToSkipSubject.setSubject(5)}
            className={`flex ${
              props.currentPage == 5
                ? `bg-indigo-400 text-white`
                : `bg-white text-gray-900`
            }  items-center justify-center px-3 h-8 leading-tight   border border-gray-300 hover:bg-gray-100 hover:text-gray-700`}
          >
            5
          </button>
          <button
            onClick={() => {
              pageToSkipSubject.setSubject(props.currentPage + 1)
            }}
            className={`flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-900 bg-white
                        border border-e-0 border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700 `}
          >
            Next
          </button>
        </div>
      </nav>
    </>
  )
}

export default Pagination
