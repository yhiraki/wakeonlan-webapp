import type { Target } from '../types';

interface DeviceCardProps {
    target: Target;
    onWake: (mac: string) => Promise<void>;
    isWaking: boolean;
}

export function DeviceCard({ target, onWake, isWaking }: DeviceCardProps) {
    return (
        <div className="card mb-1" style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
            <div>
                <h3 style={{ fontSize: 'var(--font-size-lg)', marginBottom: '0.25rem' }}>{target.Name}</h3>
                <code style={{ color: 'var(--color-text-secondary)', fontSize: 'var(--font-size-sm)' }}>
                    {target.MAC}
                </code>
            </div>
            <button
                onClick={() => onWake(target.MAC)}
                disabled={isWaking}
                style={{
                    backgroundColor: isWaking ? 'var(--color-bg-primary)' : 'var(--color-accent)',
                    color: isWaking ? 'var(--color-text-secondary)' : '#fff',
                    padding: 'var(--space-xs) var(--space-sm)',
                    borderRadius: 'var(--radius-md)',
                    fontWeight: 'bold',
                    opacity: isWaking ? 0.7 : 1,
                    transition: 'all 0.2s ease',
                }}
            >
                {isWaking ? 'Sending...' : 'Wake'}
            </button>
        </div>
    );
}
