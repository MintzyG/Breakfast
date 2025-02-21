import PancakeNote from '../../components/ui/PancakeNote';
import PancakeModal from '../../components/ui/PancakeModal';
import { useState, useEffect } from 'react';
import api from '../../../axios';

const Pancake = () => {
  const [isModalOpen, setModalOpen] = useState(false);
  const [selectedPancake, setSelectedPancake] = useState(null);
  const [pancakes, setPancakes] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPancakes = async () => {
      try {
        setLoading(true);
        const response = await api.get('/pancake');
        setPancakes(response.data);
        setError(null);
      } catch (err) {
        console.error("Error fetching pancakes:", err);
        setError("Failed to load your notes. Please try again later.");
      } finally {
        setLoading(false);
      }
    };

    fetchPancakes();
  }, []);

  const handleNoteClick = (pancake) => {
    setSelectedPancake(pancake);
    setModalOpen(true);
  };

  const handleCloseModal = () => {
    setModalOpen(false);
  };

  const handleUpdatePancake = (updatedPancake) => {
    setPancakes(prevPancakes => 
      prevPancakes.map(pancake => 
        pancake.note_id === updatedPancake.note_id ? updatedPancake : pancake
      )
    );
    setSelectedPancake(updatedPancake);
  };

  if (loading) {
    return <div className="p-4 text-center">Loading your notes...</div>;
  }

  if (error) {
    return <div className="p-4 text-center text-red-500">{error}</div>;
  }

  return (
    <>
      <div className='flex flex-row gap-4 flex-wrap p-[10px]'>
        {pancakes.length > 0 ? (
          pancakes.map((pancake) => (
            <PancakeNote
              key={pancake.note_id}
              pancake={pancake}
              onNoteClick={handleNoteClick}
            />
          ))
        ) : (
          <div className="p-4 text-center w-full">No notes found. Create your first note!</div>
        )}
      </div>
      <PancakeModal
        isOpen={isModalOpen}
        onClose={handleCloseModal}
        pancake={selectedPancake}
        onUpdate={handleUpdatePancake}
      />
    </>
  );
};

export default Pancake;
