import React from "react";
import "./Logo.css";

import logo from "../../../../images/world.svg";

export function Logo() {
    return (
        <div className="login-logo">
            <img src={logo} alt="logo" />
            <span> TravelGo </span>
        </div>
    );
}
