import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons'
import { faGithub, faInstagram, faTwitter } from '@fortawesome/free-brands-svg-icons'
import { useRef } from 'react'

function SocialMedia({ icon }) {
  return <FontAwesomeIcon icon={icon} className="relative mr-4 last:m-0 cursor-pointer transition hover:text-amber-200 hover:transition outline-indigo-400 outline-3 active:outline-dashed rounded" size="md" color="#ec4899" />;
}

export default function Header() {
  const inputRef = useRef();

  function onSearchIconClick() {
    inputRef.current.focus();
  }

  return (
    <>
    <div className="bg-zinc-900 mx-auto text-center w-full">
      <div className="absolute top-4 right-6 float-right">
          <SocialMedia icon={faTwitter} />
          <SocialMedia icon={faInstagram} />
          <SocialMedia icon={faGithub} />
      </div>
      <div className="relative inline-block mb-8 pt-20">
        <h1 className="cursor-default relative block text-6xl font-extrabold text-purple-400 font-itim drop-shadow-lg">
          Ucuz Kitapp
        </h1>
        <h4 className="cursor-default relative block border-4 border-white float-right text-left font-black font-itim text-gray-100 drop-shadow-sm transform-gpu rotate-3 p-1.5 rounded-tr-2xl rounded-br-2xl rounded-bl-2xl bg-pink-500 -bottom-2 text-md">
          💀 Ölücülerin adresi
        </h4>
      </div>
      <div className="relative mx-auto block">
        <input className="inline-block w-6/12 border-4 border-purple-400 rounded-full font-itim text-left py-4 pl-6 pr-10 mb-6 text-2xl bg-zinc-800 shadow-inner-xl outline-dashed outline-none focus:outline-indigo-400" type="text" placeholder="1984, Dijital Kale, Kayıp Tanrılar Ülkesi..." ref={inputRef} />
        <FontAwesomeIcon icon={faMagnifyingGlass} className="bg-purple-400 shadow-xl rounded-full p-4 absolute -mx-10 my-8 cursor-pointer transition hover:bg-amber-200 active:outline-dashed active:outline-indigo-400" size="xl" color="#18181b" onClick={onSearchIconClick} />
      </div>
    </div>
    <svg className="z-0 block" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1440 320"><path fill="#18181b" fill-opacity="1" d="M0,64L8.9,106.7C17.8,149,36,235,53,250.7C71.1,267,89,213,107,176C124.4,139,142,117,160,101.3C177.8,85,196,75,213,80C231.1,85,249,107,267,138.7C284.4,171,302,213,320,234.7C337.8,256,356,256,373,256C391.1,256,409,256,427,213.3C444.4,171,462,85,480,96C497.8,107,516,213,533,256C551.1,299,569,277,587,240C604.4,203,622,149,640,154.7C657.8,160,676,224,693,218.7C711.1,213,729,139,747,101.3C764.4,64,782,64,800,96C817.8,128,836,192,853,202.7C871.1,213,889,171,907,138.7C924.4,107,942,85,960,85.3C977.8,85,996,107,1013,128C1031.1,149,1049,171,1067,181.3C1084.4,192,1102,192,1120,192C1137.8,192,1156,192,1173,181.3C1191.1,171,1209,149,1227,133.3C1244.4,117,1262,107,1280,117.3C1297.8,128,1316,160,1333,165.3C1351.1,171,1369,149,1387,138.7C1404.4,128,1422,128,1431,128L1440,128L1440,0L1431.1,0C1422.2,0,1404,0,1387,0C1368.9,0,1351,0,1333,0C1315.6,0,1298,0,1280,0C1262.2,0,1244,0,1227,0C1208.9,0,1191,0,1173,0C1155.6,0,1138,0,1120,0C1102.2,0,1084,0,1067,0C1048.9,0,1031,0,1013,0C995.6,0,978,0,960,0C942.2,0,924,0,907,0C888.9,0,871,0,853,0C835.6,0,818,0,800,0C782.2,0,764,0,747,0C728.9,0,711,0,693,0C675.6,0,658,0,640,0C622.2,0,604,0,587,0C568.9,0,551,0,533,0C515.6,0,498,0,480,0C462.2,0,444,0,427,0C408.9,0,391,0,373,0C355.6,0,338,0,320,0C302.2,0,284,0,267,0C248.9,0,231,0,213,0C195.6,0,178,0,160,0C142.2,0,124,0,107,0C88.9,0,71,0,53,0C35.6,0,18,0,9,0L0,0Z"></path></svg>
    </>
  )
}
