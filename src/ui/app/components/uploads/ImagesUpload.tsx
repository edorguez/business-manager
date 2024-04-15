'use client';

import Image from "next/image";
import { ChangeEvent, useState } from "react";

const ImagesUpload = () => {
 const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  // Declare state to hold the image URLs for preview
  const [imagePreviewUrls, setImagePreviewUrls] = useState<string[]>([]);

  // Event handler for file input change
  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    const files = event.target.files; // Get the files selected by the user
    if (files) {
      // Convert files to an array
      const filesArray = Array.from(files);
      // Set the uploaded files
      setUploadedFiles(filesArray);
      // Read and set preview URLs for each file
      filesArray.forEach(file => {
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onloadend = () => {
          // Once the file is read, create a Blob URL for the image preview
          if (reader.result) {
            const blobUrl = URL.createObjectURL(file);
            setImagePreviewUrls(prevUrls => [...prevUrls, blobUrl]);
          }
        };
      });
    }
  };

  // Function to handle removing an uploaded image
  const handleRemoveImage = (index: number) => {
    const updatedFiles = [...uploadedFiles];
    const updatedPreviewUrls = [...imagePreviewUrls];

    updatedFiles.splice(index, 1); // Remove the file from the array
    updatedPreviewUrls.splice(index, 1); // Remove the preview URL from the array

    setUploadedFiles(updatedFiles);
    setImagePreviewUrls(updatedPreviewUrls);
  };

  return (
    <div>
      <input type="file" multiple onChange={handleFileChange} />
      {uploadedFiles.length > 0 && (
        <div>
          {uploadedFiles.map((file, index) => (
            <div key={index} className="relative">
              <button
                onClick={() => handleRemoveImage(index)}
                className="absolute top-5 right-5"
              >
                Remove
              </button>
              {imagePreviewUrls[index] && (
                <Image src={imagePreviewUrls[index]} alt="" width={100} height={100} />
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default ImagesUpload;
