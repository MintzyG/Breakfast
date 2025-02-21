import PropTypes from 'prop-types'

const PancakeNote = ({ pancake, onNoteClick }) => {
  const { noteId, emoji, title, content, color } = pancake;
  return (
    <div 
      className="NoteContainer w-[150px] h-[150px] cursor-pointer transition-transform hover:scale-105 rounded-xl shadow-md overflow-hidden"
      onClick={() => onNoteClick(pancake)}
      key={noteId}>
      <div className="NoteHeader flex flex-row items-center gap-1 pt-1/2 pb-1/2"
        style={{ backgroundColor: `#${color || '4CAF50'}` }}>
        <div className="EmojiSection text-md ml-2">
          { emoji || '📝' }
        </div>
        <div className="TitleSection truncate text-sm pr-1">
          { title || 'Default Title BIG ded' }
        </div>
      </div>
      <div className="NoteBody overflow-hidden text-sm pl-2 pr-2 pb-2">
        <p className="line-clamp-6">{ content || 'Default Content aaaaaaaaaaa aaaaaaaaa aaaaaaaaaa aaaaaaaaa aaaaaa aaaaaaaaaa aaaaaaaa aaaaaa aaaaa baby' }</p>
      </div>
    </div>
  )
}

PancakeNote.propTypes = {
  pancake: PropTypes.shape({
    noteId: PropTypes.number.isRequired,
    userId: PropTypes.string,
    emoji: PropTypes.string,
    title: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
    color: PropTypes.string,
    createdAt: PropTypes.string,
    updatedAt: PropTypes.string
  }).isRequired,
  onNoteClick: PropTypes.func.isRequired
};

export default PancakeNote
