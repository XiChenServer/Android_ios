package com.example.mine.fragment.listviewfragment.entity;

import java.io.File;

/**
 * @Author winiymissl
 * @Date 2024-02-24 21:06
 * @Version 1.0
 */
public class PicRecyclerViewEntity {
    File file;

    public PicRecyclerViewEntity(File file) {
        this.file = file;
    }

    public File getFile() {
        return file;
    }

    public void setFile(File file) {
        this.file = file;
    }
}
