import "./Header.css"
import { Link } from "react-router-dom";

type HeaderProps = {
    isLoggedIn: boolean;
};

export default function Header({ isLoggedIn }: HeaderProps) {
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
                        <div className="header-container_item">
                            <span>Logout</span>
                        </div>
                    )}
                    </div>
                </div>
            </div>
            <h1>Welcome, try to sign in</h1>
        </>
    )

}