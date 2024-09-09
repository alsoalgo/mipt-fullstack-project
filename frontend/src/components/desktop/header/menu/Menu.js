import "./Menu.css";
import React, { useEffect, useState } from "react";
import Button from "./button/Button";
import { Link } from 'react-router-dom';
import { ApiQuery } from "../../../../services/Api";

import help from "../../../images/help_chat.svg";
import profile from "../../../images/person-svgrepo-com.svg";

function QueryProfile(setFirstName) {
    ApiQuery('/profile', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
    }
    ).then((resp) => {
        if (resp.status != "failed") {
            var profile = resp.data.info;
            if (profile.firstName.length > 0) {
                setFirstName(profile.firstName);
            }
        }
    });
}

function Menu() {
    const [firstname, setFirstName] = useState('Кто-то');

    useEffect(() => {
        QueryProfile(setFirstName);
    }, [])

    return (
        <div className="header-menu">
            <Link to="/support" className="header-menu-button"><Button text="Поддержка" iconUrl={help}/></Link>
            <Link to="/profile" className="header-menu-button"><Button text={firstname} iconUrl={profile}/></Link>
        </div>
    );
}

export default Menu;