import { useEffect, useState } from 'react';
import { TargetList } from './components/TargetList';

function App() {
  const [version, setVersion] = useState<string>('');

  useEffect(() => {
    fetch('/api/version')
      .then(res => res.json())
      .then(data => setVersion(data.version))
      .catch(err => console.error('Failed to fetch version:', err));
  }, []);

  return (
    <div className="container">
      <header className="text-center mb-2">
        <h1>Wake on LAN</h1>
        <p style={{ color: 'var(--color-text-secondary)' }}>
          Select a device to wake up
        </p>
      </header>

      <main>
        <TargetList />
      </main>

      <footer className="text-center mt-2">
        <p style={{ fontSize: 'var(--font-size-sm)', color: 'var(--color-text-secondary)' }}>
          &copy; 2026 Wake on LAN Web <span style={{ opacity: 0.7 }}>{version && `(${version})`}</span>
        </p>
      </footer>
    </div>
  )
}

export default App
