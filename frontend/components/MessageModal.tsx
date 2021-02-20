import React, { useEffect, useRef, useState } from "react"
import styled from "styled-components"

type Props = {
    message: string;
    visible: boolean;
    onClose: Function;
}

export default function MessageModal({ message, visible, onClose }: Props) {
    // const [visible, onClose] = useState(false)
    const wrapperRef = useRef(null)
    userCLickedOutsideHandler(wrapperRef, onClose)

    return (
        <Modal ref={wrapperRef} style={{ display: visible ? "block" : "none" }}>

            <ModalContent>
                <Close onClick={() => onClose()}>&times;</Close>
                <h1>{message}</h1>
            </ModalContent>

        </Modal>
    )
}

function userCLickedOutsideHandler(ref, onClose) {
    useEffect(() => {
        /**
         * Alert if clicked on outside of element
         */
        function handleClickOutside(event) {
            if (ref.current && !ref.current.contains(event.target)) {
                onClose()
            }
        }

        // Bind the event listener
        // document.addEventListener("mousedown", handleClickOutside);
        // return () => {
        //     // Unbind the event listener on clean up
        //     document.removeEventListener("mousedown", handleClickOutside);
        // };
    }, [ref]);
}

/* The Modal (background) */
const Modal = styled.div`
    display: none; /* Hidden by default */
    position: fixed; /* Stay in place */
    z-index: 1; /* Sit on top */
    left: 0;
    top: 0;
    width: 100%; /* Full width */
    height: 100%; /* Full height */
    overflow: auto; /* Enable scroll if needed */
    //background-color: rgb(0,0,0); /* Fallback color */
    background-color: rgba(0,0,0,0.4); /* Black w/ opacity */
    color: #00AA00;
`

/* Modal Content/Box */
const ModalContent = styled.div`
    background-color: #fefefe;
    margin: 15% auto; /* 15% from the top and centered */
    padding: 20px;
    border: 1px solid #888;
    width: 80%; /* Could be more or less, depending on screen size */
  `

/* The Close Button */
const Close = styled.span`
    color: #aaa;
    float: right;
    font-size: 28px;
    font-weight: bold;

    :hover,
    :focus {
        color: black;
        text-decoration: none;
        cursor: pointer;
    }
`