import {useState} from 'react';
import logo from './assets/images/logo-minecraft.png';
import './App.css';
import ToggleSwitchUI from "./components/UI/ToggleSwitchUI.jsx";
import NeonBeamContainer from "./components/NeonBeamContainer.jsx";
import Logo from "./components/Logo.jsx";
import {ConnectSSH, DisconnectSSH} from "../wailsjs/go/main/App.js";
import ErrorContainer from "./components/ErrorContainer.jsx";
import { EventsOn } from '@wailsapp/runtime';

function App() {
    const [isOn, setIsOn] = useState(false);
    const [isLoad, setIsLoad] = useState(false);
    const [error, setError] = useState(null);

    const updateResultSSH = (result) => {
        if (result.status)  {
            setIsOn(!isOn);
        } else {
            setError(result.meta);
        }
        setIsLoad(false);
    };

    function connection() {
        setError(null);
        setIsLoad(true);
        if (!isOn) {
            ConnectSSH().then(updateResultSSH);
        } else {
            DisconnectSSH().then(updateResultSSH);
        }
    }


    window.runtime.EventsOn("connectionEndError", (message) => {
        setError(message);
        setIsOn(false);
        setIsLoad(false);
    });

    window.runtime.EventsOn("connectionEnd", () => {
        setIsOn(false);
        setIsLoad(false);
    });

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo"/>
            <Logo/>
            <div className={'header'}>
                <div className={'header__top-text'}>Для подключения используйте адрес сервера:</div>
                <div className={'header__text'}>localhost</div>
                <div className={'header__sub-text'}> При подключении будет создан ssh-тунель с сервером Minecraft на локальный порт 25565</div>
            </div>
            <NeonBeamContainer status={isOn}>
                <ToggleSwitchUI
                    isOn={isOn}
                    isLoading={isLoad}
                    handleToggle={connection}
                    onColor="#04B86C"
                    offColor="#7d7d7d"
                />
            </NeonBeamContainer>
            <ErrorContainer error={error} />
        </div>
    )
}

export default App
