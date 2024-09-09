import "./Bar.css";
import React, { useEffect, useState } from "react";
import { useNavigate, createSearchParams } from 'react-router-dom';

export function Bar({city, from, to, isHotelPressed}) {
    const navigate = useNavigate();
    const [hotelWhere, setHotelWhere] = useState('');
    const [hotelDateIn, setDateIn] = useState('');
    const [hotelDateOut, setDateOut] = useState('');

    useEffect(() => {
        setHotelWhere(city);
        setDateIn(from);
        setDateOut(to);
    }, [])

    const handleHotelWhereChange = (event) => {
        setHotelWhere(event.target.value);
    }

    const handleHotelDateInChange = (event) => {
        setDateIn(event.target.value);
    }

    const handleHotelDateOutChange = (event) => {
        setDateOut(event.target.value);
    }

    const search = async (event) => {
        event.preventDefault()

        const path = {
            pathname: '/hotels',
            search: createSearchParams({
                "city": hotelWhere,
                "from": hotelDateIn,
                "to": hotelDateOut,
            }).toString()
        };
        navigate(path);
    }

    const onFocusHandler = (event) => {
        event.target.type = 'date';
    }

    return (
        <div key="search-bar-hotel" className="search-bar">
            <input key="search-bar-input-where" type="text" className="search-bar-input" placeholder="Куда хотите поехать?" value={hotelWhere} onChange={handleHotelWhereChange}/>
            <div className="search-bar-vertical-line"></div>
            <input key="search-bar-input-from" type="text" className="search-bar-date-input" placeholder="Заезд" value={hotelDateIn} onFocus={onFocusHandler} onChange={handleHotelDateInChange}/>
            <div className="search-bar-vertical-line"></div>
            <input key="search-bar-input-to" type="text" className="search-bar-date-input" placeholder="Выезд" value={hotelDateOut} onFocus={onFocusHandler} onChange={handleHotelDateOutChange}/>
            <div className="search-bar-vertical-line"></div>
            <button className="search-bar-button" onClick={search}>Найти</button>
        </div>
    );
}
