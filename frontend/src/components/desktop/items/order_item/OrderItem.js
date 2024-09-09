import './OrderItem.css';
import React from 'react';

export function OrderItem({title, description, img}) {
    return (
        <div className='order-item'>
            <img className='order-item-bg' alt='bg' src={img}></img>
            <div className='order-item-description-wrapper'>
                <h1 className='order-item-description-title'>{title}</h1>
                <p className='order-item-description-sub'>{description}</p>
            </div>
        </div>
    )
}