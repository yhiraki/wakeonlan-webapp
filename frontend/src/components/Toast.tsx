import { useEffect, useState } from 'react';

export type ToastType = 'success' | 'error';

export interface ToastProps {
    message: string;
    type: ToastType;
    show: boolean;
    onClose: () => void;
}

export function Toast({ message, type, show, onClose }: ToastProps) {
    const [visible, setVisible] = useState(false);

    useEffect(() => {
        if (show) {
            setVisible(true);
            const timer = setTimeout(() => {
                setVisible(false);
                setTimeout(onClose, 300); // Wait for animation
            }, 3000);
            return () => clearTimeout(timer);
        }
    }, [show, onClose]);

    if (!show && !visible) return null;

    return (
        <div className={`toast toast-${type} ${visible ? 'show' : ''}`}>
            <div className="toast-content">{message}</div>
        </div>
    );
}
