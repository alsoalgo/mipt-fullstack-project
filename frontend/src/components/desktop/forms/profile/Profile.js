import './Profile.css';
import React, { useEffect, useState } from 'react';
import { ApiQuery } from '../../../../services/Api';

function ChangeProfile(firstname, lastname, surname) {
    ApiQuery('/profile/edit', {
        method: 'post',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            "firstName": firstname,
            "lastName": lastname,
            "surName": surname,
        })
    }
    ).then((resp) => {});
} 

function QueryProfile(setFirstName, setLastName, setSurName) {
    ApiQuery('/profile', {
        method: 'get',
        headers: {
            'Content-Type': 'application/json',
        },
    }
    ).then((resp) => {
        var profile = resp.data.info;
        setFirstName(profile.firstName);
        setLastName(profile.lastName);
        setSurName(profile.surName);
    });
}

export function Form() {
    const [firstname, setFirstName] = useState('');
    const [lastname, setLastName] = useState('');
    const [surname, setSurName] = useState('');

    useEffect(() => {
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

        ChangeProfile(firstname, lastname, surname);
    }

    return (
        <div className="profile-edit-form">
            <h1>Обо мне</h1>
            <input key="profile-edit-form-input-firstname" className='profile-edit-form-input' placeholder='Имя' onChange={handleFirstNameChange} value={firstname} ></input>
            <input key="profile-edit-form-input-lastname" className='profile-edit-form-input' placeholder='Фамилия' onChange={handleLastNameChange} value={lastname}></input>
            <input key="profile-edit-form-input-surname" className='profile-edit-form-input' placeholder='Отчество' onChange={handleSurNameChange} value={surname}></input>
            <button className='profile-edit-form-submit' onClick={submitForm}>Исправить</button>
        </div>
    );
}