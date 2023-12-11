import Modal from 'react-modal';
import React, {ReactNode} from "react";

interface CustomModalProps {
    isOpen: boolean;
    onRequestClose: () => void;
    children: ReactNode;
    label: string;
}

export const CustomModal: React.FC<CustomModalProps> = ({isOpen, onRequestClose, children, label}) => {
    return (
        <Modal
            className=" overflow-y-auto  w-3/4 fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 bg-white p-4 rounded-xl shadow-xl h-4/5 "
            overlayClassName="fixed inset-0 bg-gray-700 bg-opacity-50"
            isOpen={isOpen}
            onRequestClose={onRequestClose}
            contentLabel={label}
        >
            {children}
        </Modal>
    );
};
