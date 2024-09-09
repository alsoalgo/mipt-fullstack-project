import './Order.css';
import React, { useEffect, useState } from 'react';
import { ApiQuery, TokenCheck } from '../../../../services/Api';
import { useNavigate, useSearchParams } from 'react-router-dom';
        
function QueryProfile(setFirstName, setLastName, setSurName) {
    ApiQuery('/profile', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
    }
    ).then((resp) => {
        if (resp.status != "failed") {
            var profile = resp.data.info;
            setFirstName(profile.firstName);
            setLastName(profile.lastName);
            setSurName(profile.surName);
        }
    });
}

function OrderHotel(hotel, person, navigate) {
    ApiQuery('/order', {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            hotelId: Number(hotel.hotelId),
            dateFrom: hotel.from,
            dateTo: hotel.to,
            firstName: person.firstname,
            lastName: person.lastname,
            surName: person.surname,
        })
    }
    ).then((resp) => {
        if (resp.status != "failed") {
            navigate("/profile");
        }
    });
}

export function Form() {
    let navigate = useNavigate();
    const [searchParams, setSearchParams] = useSearchParams();
    var [hotelId, setHotelId] = useState(searchParams.get("hotelId"));
    var [from, setFrom] = useState(searchParams.get("from"));
    var [to, setTo] = useState(searchParams.get("to"));

    const [firstname, setFirstName] = useState('');
    const [lastname, setLastName] = useState('');
    const [surname, setSurName] = useState('');

    useEffect(() => {
        TokenCheck().then((ok) => {
            if (!ok) {
                navigate('/login');
            }
        });

        QueryProfile(setFirstName, setLastName, setSurName);
    }, []);

    const handleFirstNameChange = (event) => {
        setFirstName(event.target.value);
    }

    const handleLastNameChange = (event) => {
        setLastName(event.target.value);
    }

    const handleSurNameChange = (event) => {
        setSurName(event.target.value);
    }

    const submitForm = (event) => {
        event.preventDefault();
        
        OrderHotel(
            {hotelId, from, to}, 
            {firstname, lastname, surname},
            navigate
        );
    }

    return (
        <div className="order-form">
            <h1>Информация о покупателе</h1>
            <input className='order-form-passenger-input' placeholder='Имя' onChange={handleFirstNameChange} value={firstname} ></input>
            <input className='order-form-passenger-input' placeholder='Фамилия' onChange={handleLastNameChange} value={lastname}></input>
            <input className='order-form-passenger-input' placeholder='Отчество' onChange={handleSurNameChange} value={surname}></input>
            <button className='order-form-submit' onClick={submitForm}>Забронировать</button>
        </div>
    );
}