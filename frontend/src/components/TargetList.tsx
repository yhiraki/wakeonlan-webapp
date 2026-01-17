import { useEffect, useState } from 'react';
import type { Target } from '../types';
import { DeviceCard } from './DeviceCard';
import { Toast, type ToastType } from './Toast';

export function TargetList() {
    const [targets, setTargets] = useState<Target[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);
    const [wakingMac, setWakingMac] = useState<string | null>(null);

    // Toast State
    const [toast, setToast] = useState<{ show: boolean; message: string; type: ToastType }>({
        show: false,
        message: '',
        type: 'success',
    });

    const showToast = (message: string, type: ToastType) => {
        setToast({ show: true, message, type });
    };

    const closeToast = () => {
        setToast(prev => ({ ...prev, show: false }));
    };

    useEffect(() => {
        fetchTargets();
    }, []);

    const fetchTargets = async () => {
        try {
            const res = await fetch('/api/targets');
            if (!res.ok) throw new Error('Failed to fetch targets');
            const data = await res.json();
            setTargets(data);
        } catch (err) {
            // For development/demo without backend running on same port, we might want to mock if fetch fails 
            // strict implementation: show error
            setError('Could not load targets. Is the server running?');
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleWake = async (mac: string) => {
        setWakingMac(mac);
        try {
            const res = await fetch('/api/wake', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ mac }),
            });

            if (!res.ok) {
                const errData = await res.json().catch(() => ({}));
                throw new Error(errData.error || 'Failed to send Wake command');
            }

            showToast(`Magic Packet sent to ${mac}`, 'success');

        } catch (err: any) {
            console.error(err);
            showToast(err.message || 'Failed to wake device', 'error');
        } finally {
            // Simulate a small delay for better UX if response is too fast
            setTimeout(() => setWakingMac(null), 500);
        }
    };

    if (loading) {
        return (
            <div className="card text-center" style={{ padding: 'var(--space-xl)' }}>
                <p style={{ color: 'var(--color-text-secondary)' }}>Loading devices...</p>
            </div>
        );
    }

    if (error) {
        return (
            <div className="card text-center" style={{ padding: 'var(--space-xl)', borderColor: 'var(--color-error)' }}>
                <p style={{ color: 'var(--color-error)' }}>{error}</p>
                <button onClick={fetchTargets} className="mt-1" style={{ color: 'var(--color-accent)' }}>Retry</button>
            </div>
        );
    }

    if (targets.length === 0) {
        return (
            <div className="card text-center" style={{ padding: 'var(--space-xl)' }}>
                <p style={{ color: 'var(--color-text-secondary)' }}>No devices configured.</p>
                <p style={{ fontSize: 'var(--font-size-sm)', marginTop: 'var(--space-xs)', color: 'var(--color-text-secondary)' }}>
                    Start the server with <code>name=mac</code> arguments.
                </p>
            </div>
        );
    }

    return (
        <div className="target-list">
            {targets.map((target, index) => (
                <DeviceCard
                    key={`${target.MAC}-${index}`}
                    target={target}
                    onWake={handleWake}
                    isWaking={wakingMac === target.MAC}
                />
            ))}
            <Toast
                message={toast.message}
                type={toast.type}
                show={toast.show}
                onClose={closeToast}
            />
        </div>
    );
}
