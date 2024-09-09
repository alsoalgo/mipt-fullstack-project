import './PopularDestination.css';
import React, { useState } from 'react';

export function PopularDestination({title, description, img}) {

    return (
        <div className='hotels-popular-destination'>
            <img className='hotels-popular-destination-bg' alt='bg' src={img}></img>
            <h1 className='hotels-popular-destination-title'>{title}</h1>
            <p className='hotels-popular-destination-description'>{description}</p>
        </div>
    );
}