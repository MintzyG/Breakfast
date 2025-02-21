import PancakeNote from '../../components/ui/PancakeNote';

const Pancake = () => {
  const noteCount = 5;
  return (
    <div className=' flex flex-row gap-4 flex-wrap p-[10px]'>
    {Array.from({length: noteCount}).map((_, index) => (
      <PancakeNote
        key={index}
        pancake={{}}
        onNoteClick={() => console.log(`Note ${index} clicked`)}/>
    ))}
    </div>
 );
};

export default Pancake;
