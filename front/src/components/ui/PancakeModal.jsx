// PancakeModal.jsx with Color Picker implementation
import ReactDOM from "react-dom";
import PropTypes from 'prop-types';
import { PancakePropTypes } from '../../../propTypes.js';
import { X } from 'lucide-react';
import { useState, useEffect, useRef } from 'react';
import EmojiPicker from 'emoji-picker-react';
import { HexColorPicker } from 'react-colorful';

const PancakeModal = ({ isOpen, onClose, pancake, onUpdate }) => {
  // Move useState calls to the top level, before any conditionals
  const { emoji, title, content, color } = pancake || {};
  const [modalTitle, setModalTitle] = useState(title || '');
  const [modalContent, setModalContent] = useState(content || '');
  const [modalEmoji, setModalEmoji] = useState(emoji || '📝');
  const [modalColor, setModalColor] = useState(color || '4CAF50');
  const [showEmojiPicker, setShowEmojiPicker] = useState(false);
  const [showColorPicker, setShowColorPicker] = useState(false);
  
  const emojiPickerRef = useRef(null);
  const emojiContainerRef = useRef(null);
  const colorPickerRef = useRef(null);
  const colorContainerRef = useRef(null);
  
  // Update state when the prop values change (in case pancake props are updated)
  useEffect(() => {
    if (isOpen) {
      setModalTitle(title || '');
      setModalContent(content || '');
      setModalEmoji(emoji || '📝');
      setModalColor(color || '4CAF50');
    }
  }, [isOpen, title, content, emoji, color]);

  // Handle clicks outside the emoji picker to close it
  useEffect(() => {
    const handleOutsideClick = (event) => {
      // Handle emoji picker clicks
      if (
        showEmojiPicker &&
        emojiPickerRef.current && 
        !emojiPickerRef.current.contains(event.target) &&
        emojiContainerRef.current &&
        !emojiContainerRef.current.contains(event.target)
      ) {
        setShowEmojiPicker(false);
      }
      
      // Handle color picker clicks
      if (
        showColorPicker &&
        colorPickerRef.current && 
        !colorPickerRef.current.contains(event.target) &&
        colorContainerRef.current &&
        !colorContainerRef.current.contains(event.target)
      ) {
        setShowColorPicker(false);
      }
    };

    document.addEventListener('mousedown', handleOutsideClick);
    return () => {
      document.removeEventListener('mousedown', handleOutsideClick);
    };
  }, [showEmojiPicker, showColorPicker]);

  // Handle backdrop click to close the modal
  const handleBackdropClick = (event) => {
    if (event.target === event.currentTarget) {
      onClose();
    }
  };

  // Handle updates to propagate changes
  const handleTitleChange = (e) => {
    const newTitle = e.target.value;
    setModalTitle(newTitle);
    
    // Propagate changes up to parent component
    if (onUpdate && pancake) {
      onUpdate({
        ...pancake,
        title: newTitle
      });
    }
  };

  const handleContentChange = (e) => {
    const newContent = e.target.value;
    setModalContent(newContent);
    
    // Propagate changes up to parent component
    if (onUpdate && pancake) {
      onUpdate({
        ...pancake,
        content: newContent
      });
    }
  };

  const handleEmojiClick = () => {
    setShowEmojiPicker(!showEmojiPicker);
    setShowColorPicker(false);
  };

  const handleColorClick = () => {
    setShowColorPicker(!showColorPicker);
    setShowEmojiPicker(false);
  };

  const handleEmojiSelect = (emojiData) => {
    // Update the local emoji state
    setModalEmoji(emojiData.emoji);
    setShowEmojiPicker(false);
    
    // Propagate changes up to parent component
    if (onUpdate && pancake) {
      onUpdate({
        ...pancake,
        emoji: emojiData.emoji
      });
    }
  };

  const handleColorChange = (newColor) => {
    // Remove the # from the hex color
    const hexColor = newColor.replace('#', '');
    setModalColor(hexColor);
    
    // Propagate changes up to parent component
    if (onUpdate && pancake) {
      onUpdate({
        ...pancake,
        color: hexColor
      });
    }
  };

  // Return null after the hooks are defined
  if (!isOpen) return null;
  
  return ReactDOM.createPortal(
    <div
      className="fixed inset-0 flex items-center justify-center z-50 backdrop-blur-xs"
      onClick={handleBackdropClick}
    >
      <div className="relative bg-white p-6 rounded-lg shadow-lg w-[66vw] h-[95vh] min-w-sm border">
        <div className="ModalContents flex flex-col gap-2 h-full">
          {/* Modal Header */}
          <div className="ModalHeader flex flex-row items-center gap-2 py-2 w-full">
            {/* Emoji Circle */}
            <div 
              ref={emojiContainerRef}
              className="EmojiContainer w-10 h-10 flex items-center justify-center rounded-full bg-gray-300 shadow-md text-white cursor-pointer hover:bg-gray-400"
              onClick={handleEmojiClick}
            >
              <span>{modalEmoji}</span>
            </div>
            
            {/* Color Circle */}
            <div 
              ref={colorContainerRef}
              className="ColorContainer w-10 h-10 flex items-center justify-center rounded-full bg-gray-300 shadow-md cursor-pointer hover:bg-gray-400 relative"
              onClick={handleColorClick}
            >
              <div 
                className="w-6 h-6 rounded-full"
                style={{ backgroundColor: `#${modalColor}` }}
              >
                <div className="w-2 h-2 rounded-full bg-white absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"></div>
              </div>
            </div>
            
            {/* Emoji Picker */}
            {showEmojiPicker && (
              <div 
                ref={emojiPickerRef}
                className="absolute top-16 left-6 z-10"
              >
                <EmojiPicker onEmojiClick={handleEmojiSelect} />
              </div>
            )}
            
            {/* Color Picker */}
            {showColorPicker && (
              <div 
                ref={colorPickerRef}
                className="absolute top-16 left-16 z-10"
              >
                <HexColorPicker 
                  color={`#${modalColor}`} 
                  onChange={handleColorChange} 
                />
              </div>
            )}
            
            {/* Title Container with Pill Shape */}
            <div className="TitleContainer flex-1 px-4 py-2 rounded-full bg-gray-300 text-black shadow-md">
              <input
                type="text"
                className="w-full bg-transparent border-none outline-none text-black"
                value={modalTitle}
                onChange={handleTitleChange}
              />
            </div>
            {/* Close Button */}
            <button
              onClick={onClose}
              className="text-black w-10 h-10 flex justify-center items-center rounded-full bg-gray-300 hover:bg-gray-500"
            >
              <X />
            </button>
          </div>
          {/* Modal Content */}
          <div className="OuterContent flex h-full">
            <div className="ContentContainer pr-4 pt-4 pb-4 flex-1 h-full">
              <textarea
                className="w-full h-full bg-gray-300 p-2 rounded resize-none overflow-auto"
                value={modalContent}
                onChange={handleContentChange}
              />
            </div>
            {/* Options Sidebar */}
            <div className="OptionsContainer flex flex-col pt-4 w-32 space-y-2">
              {/* Option buttons */}
              <button className="w-full py-2 px-4 rounded-lg bg-gray-400 hover:bg-gray-500 text-white">
                Option 1
              </button>
              <button className="w-full py-2 px-4 rounded-lg bg-gray-400 hover:bg-gray-500 text-white">
                Option 2
              </button>
              <button className="w-full py-2 px-4 rounded-lg bg-gray-400 hover:bg-gray-500 text-white">
                Option 3
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>,
    document.getElementById("modal-root")
  );
};

PancakeModal.propTypes = {
  ...PancakePropTypes,
  isOpen: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  onUpdate: PropTypes.func
};

export default PancakeModal;
