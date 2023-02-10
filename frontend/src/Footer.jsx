import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faDiscord, faGithub, faInstagram, faTwitter } from '@fortawesome/free-brands-svg-icons'

//socialMediaStyle = "mr-4 cursor-pointer";

function SocialMediaIcon({ icon }) {
  return <FontAwesomeIcon icon={icon} className="relative mr-4 last:m-0 cursor-pointer transition hover:text-amber-200" size="lg" color="#c084fc" />;
}

export default function Footer() {
  return (
    <>
    <div className="absolute float-right top-4 right-4 font-itim text-md text-right m-auto block">
        <SocialMediaIcon icon={faTwitter} />
        <SocialMediaIcon icon={faInstagram} />
        <SocialMediaIcon icon={faGithub} />
    </div>
    </>
    );
}
