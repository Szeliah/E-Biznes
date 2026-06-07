import "./Header.css"
import { Link, useNavigate } from "react-router-dom";

const API_URL = import.meta.env.VITE_API_URL;

type HeaderProps = {
    isLoggedIn: boolean;
    setIsLoggedIn: React.Dispatch<
        React.SetStateAction<boolean>
    >;
    userEmail: string;
};

export default function Header({ isLoggedIn, setIsLoggedIn, userEmail }: HeaderProps) {
    
    const navigate = useNavigate();

    async function handleLogout() {

        await fetch(
            `${API_URL}/logout`,
            {
                method: "POST",
                credentials: "include",
            }
        );

        setIsLoggedIn(false);
        navigate("/");
    }
    
    
    return (
        <>
            <div className="header-container">
                <div className="header-container__left">
                    <div className="header-container__left-item header-container_item">
                        <span>About us</span>
                    </div>
                    <div className="header-container__left-item header-container_item">
                        <span>FAQ</span>
                    </div>
                    <div className="header-container__left-item header-container_item">
                        <span>Learn more</span>
                    </div>
                </div>
                <div className="header-container__right">
                    <div className="header-container__right-item header-container_item">
                    {!isLoggedIn ? (
                        <div className="header-container_item">
                            <Link to="/signin">
                                Sign in
                            </Link>
                        </div>
                    ) : (
                        <div className="header-container_item logged-user">
                            <span>{userEmail}</span>
                            <span onClick={handleLogout}>Logout</span>
                        </div>
                    )}
                    </div>
                </div>
            </div>
            <h1>Welcome, try to sign in</h1>
        </>
    )

}