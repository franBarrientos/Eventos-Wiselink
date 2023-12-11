import  { FunctionComponent } from 'react';

interface OwnProps {}

type Props = OwnProps;

const Footer: FunctionComponent<Props> = () => {

  return (
      <footer className="text-gray-600 body-font shadow-2xl ">
          <div className="container px-5 py-8 mx-auto flex items-center sm:flex-row flex-col">
              <a className="flex title-font font-medium items-center md:justify-start justify-center text-gray-900">
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" className="w-10 h-10 text-white p-2 bg-indigo-500 rounded-full" viewBox="0 0 24 24">
                      <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"></path>
                  </svg>
                  <span className="ml-3 text-xl">Wiselink Events</span>
              </a>
              <p className="text-lg text-gray-500 sm:ml-4 sm:pl-4 sm:border-l-2 sm:border-gray-200 sm:py-2 sm:mt-0 mt-4">© 2023 Barrientos Franco —
              </p>
              <a href="https://franservices.netlify.app/" className="text-indigo-500 ml-1 text-lg" rel="noopener noreferrer" target="_blank">franservices.netlify.app</a>
              <span className="inline-flex sm:ml-auto sm:mt-0 mt-4 justify-center sm:justify-start">
      <a href="https://www.linkedin.com/in/franco-barrientos/" className="ml-3 text-gray-500">
        <svg fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="0" className="w-5 h-5" viewBox="0 0 24 24">
          <path stroke="none" d="M16 8a6 6 0 016 6v7h-4v-7a2 2 0 00-2-2 2 2 0 00-2 2v7h-4v-7a6 6 0 016-6zM2 9h4v12H2z"></path>
          <circle cx="4" cy="4" r="2" stroke="none"></circle>
        </svg>
      </a>
    </span>
          </div>
      </footer>
  );
};

export default Footer;
