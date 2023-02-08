import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons'
import { useRef } from 'react'

export default function Header() {
  const inputRef = useRef();

  function onSearchIconClick() {
    inputRef.current.focus();
  }

  return (
    <>
    <div className="bg-zinc-900 mx-auto text-center">
      <div className="relative inline-block mb-8 pt-12">
        <h1 className="relative block text-6xl font-extrabold text-purple-400 font-zilla drop-shadow-lg">
          Ucuz Kitapp
        </h1>
        <h4 className="relative block border-4 border-white float-right text-left font-black font-zilla text-gray-100 drop-shadow-sm transform-gpu rotate-3 p-1.5 rounded-tr-2xl rounded-br-2xl rounded-bl-2xl bg-pink-500 -bottom-2">
          💀 Ölücülerin adresi
        </h4>
      </div>

      <div className="relative mx-auto block">
        <input className="inline-block w-6/12 border-4 border-purple-400 rounded-full font-zilla text-left py-4 px-6 mb-6 text-xl bg-zinc-800 shadow-inner-xl outline-dashed outline-none focus:outline-indigo-400" type="text" placeholder="1984, Dijital Kale, Kayıp Tanrılar Ülkesi..." ref={inputRef} />
        <FontAwesomeIcon icon={faMagnifyingGlass} className="bg-purple-400 shadow-xl rounded-full p-4 absolute -mx-10 my-8 cursor-pointer transition hover:bg-amber-200" size="xl" color="#18181b" onClick={onSearchIconClick} />
      </div>
    </div>
    <div className="wave bg-zinc-900"></div>
    </>
  )
}
