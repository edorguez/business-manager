'use client';

import { Icon } from "@iconify/react";
import Image from "next/image";
import { ChangeEvent, useState } from "react";

const ImagesUpload = () => {
  const [uploadedFiles, setUploadedFiles] = useState<File[]>([]);
  const [imagePreviewUrls, setImagePreviewUrls] = useState<string[]>([]);

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
    <>
      <div className="flex justify-center mb-5">
        <input type="file" multiple onChange={handleFileChange} className="
          block 
          text-sm 
          text-slate-500
          file:mr-4 
          file:py-2 
          file:px-4
          file:rounded-full 
          file:border-0
          cursor-pointer
          file:text-sm 
          file:font-semibold
          file:bg-maincolorhov
          file:text-maincolor
          hover:file:bg-thirdcolorhov
          hover:file:text-thirdcolor
          file:transition
          file:duration-150" 
        />
      </div>

      <hr />

      {uploadedFiles.length > 0 && (
        <div className="mt-6 grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 place-items-center">
          {uploadedFiles.map((file, index) => (
            <div key={index} className="relative w-[100px]">
              <button
                onClick={() => handleRemoveImage(index)}
                className="absolute top-[-10px] right-[-10px] rounded-full bg-thirdcolor hover:bg-rose-600 text-white text-lg p-1 transition duration-100"
              >
                <Icon icon="material-symbols:close" />
              </button>
              {imagePreviewUrls[index] && (
                <Image src={imagePreviewUrls[index]} alt="" width={100} height={100} />
              )}
            </div>
          ))}
        </div>
      )}

      {uploadedFiles.length === 0 && (
        <span>Sube</span>
      )}
    </>
  );
}

export default ImagesUpload;
