import React from 'react';

function App() {
  return (
    <div className="min-h-screen w-screen p-5 bg-gradient-to-r from-indigo-500 via-purple-500 to-pink-500">
      <Card>
        <h1 className="text-4xl font-black">
          Alex Stojda
        </h1>
      </Card>
    </div>
  );
}

function Card({children}: { children: React.ReactNode }) {
  return (
    <div className="container mx-auto bg-white/30 p-5 rounded-lg drop-shadow-md">
      {children}
    </div>
  )
}

export default App;
