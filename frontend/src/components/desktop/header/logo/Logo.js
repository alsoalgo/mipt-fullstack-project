import React from "react";
import "./Logo.css";

import logo from "../../../images/world.svg";

function Logo() {
    return (
        <div className="header-logo">
            <img src={logo} alt="logo" />
            <span> TravelGo </span>
        </div>
    );
}

export default Logo;