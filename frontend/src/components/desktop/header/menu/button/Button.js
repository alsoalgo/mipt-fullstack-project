import React from "react";
import "./Button.css";

function Button({text, iconUrl}) {
    return (
        <button className="header-menu-button-inner">
            <img src={iconUrl}></img> 
            <span>{text}</span>
        </button>
    );
}

export default Button;