export default function Header() {
  return (
    <div className="bg-zinc-900 mx-auto text-center">
      <div className="relative inline-block mb-5 pt-12">
        <h1 className="relative block text-6xl font-extrabold text-purple-400 font-zilla drop-shadow-lg">
          Ucuz Kitapp
        </h1>
        <h4 className="relative block text-left bottom-2 font-zilla text-amber-200 drop-shadow-sm">
          💰 Ölücülerin adresi.
        </h4>
      </div>

      <div className="mx-auto block">
        <input className="inline-block w-6/12 border-4 border-purple-400 rounded-full text-left py-4 px-6 text-xl bg-zinc-800 shadow-inner outline-none" type="text" placeholder="1984, Dijital Kale, Kayıp Tanrılar Ülkesi..." />
      </div>
    </div>
  )
}
