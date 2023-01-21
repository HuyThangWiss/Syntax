//
// Created by huythang on 1/21/23.
//

#ifndef P1_SINHVIEN_H
#define P1_SINHVIEN_H
#include <string>
#include <iostream>
#include <vector>
#include <fstream>
#include <sys/stat.h>
#include <sstream>

using namespace std;

class sinhVien {
private:
    sinhVien(string basicString, string basicString1, string basicString2, string basicString3);

    string ten;
    int tuoi;
    string diachi;
    string sdt;

public:
    string getTen() const { return ten; }
    void setTen(string ten) { this->ten = ten; }

    int getTuoi() const { return tuoi; }
    void setTuoi(int tuoi) { this->tuoi = tuoi; }

    string getDiachi() const { return diachi; }
    void setDiachi(string diachi) { this->diachi = diachi; }

    string getSdt() const { return sdt; }
    void setSdt(string sdt) { this->sdt = sdt; }


    sinhVien();

    vector<sinhVien> dsThongTin;
    void nhap();
    void hienThi();
    void themThongTin();
    void timKiemTheoTen(string ten);
    void xoaThongTinTheoTen(string ten);
    void themThongTinSauTen(string ten);
    void ghiDuLieuVaoFile();
    void docDuLieuTuFile();


};
void sinhVien::nhap() {
    cout << "Nhap ten: ";
    cin >> ten;
    cout << "Nhap tuoi: ";
    cin >> tuoi;
    cout << "Nhap dia chi: ";
    cin >> diachi;
    cout << "Nhap so dien thoai: ";
    cin >> sdt;
}

void sinhVien::hienThi(){
    cout << "Ten: " << ten << endl;
    cout << "Tuoi: " << tuoi << endl;
    cout << "Dia chi: " << diachi << endl;
    cout << "So dien thoai: " << sdt << endl;
}
void sinhVien::themThongTin(){
    int n;
    cout<<"Nhap so luong sinh vien"<<endl;
    cin>>n;
    sinhVien t;
    for(int i=0;i<n;i++){
        cout<<"Nhap sinh vien thu "<<i+1<<endl;
        t.nhap();
        dsThongTin.push_back(t);
    }


}

void sinhVien::timKiemTheoTen(string ten) {
    for (int i = 0; i < dsThongTin.size(); i++) {
        if (dsThongTin[i].getTen() == ten) {
            dsThongTin[i].hienThi();
        }
    }
}
void sinhVien::xoaThongTinTheoTen(string ten) {
    for (int i = 0; i < dsThongTin.size(); i++) {
        if (dsThongTin[i].getTen() == ten) {
            dsThongTin.erase(dsThongTin.begin() + i);
            break;
        }
    }
}
void sinhVien::themThongTinSauTen(string ten) {
    for (int i = 0; i < dsThongTin.size(); i++) {
        if (dsThongTin[i].getTen() == ten) {
            sinhVien t;
            t.nhap();
            dsThongTin.insert(dsThongTin.begin() + i + 1, t);
            break;
        }
    }
}
void sinhVien::ghiDuLieuVaoFile() {
    struct stat buffer;
    if (stat("SinhVien.txt", &buffer) != 0) {
        ofstream file("SinhVien.txt");
        file.close();
    }

    ofstream file("SinhVien.txt", ios::app);
    for (int i = 0; i < dsThongTin.size(); i++) {
        bool tonTai = false;
        ifstream fileIn("SinhVien.txt");
        string line;
        while (getline(fileIn, line)) {
            if (line.find(dsThongTin[i].getTen()) != string::npos) {
                tonTai = true;
                break;
            }
        }
        fileIn.close();
        if(!tonTai) {
            file << dsThongTin[i].getTen() << " ";
            file << dsThongTin[i].getTuoi() << " ";
            file << dsThongTin[i].getDiachi() << " ";
            file << dsThongTin[i].getSdt() << endl;
        }
    }
    file.close();
}
void sinhVien::docDuLieuTuFile() {
    ifstream file("SinhVien.txt");
    string line;
    while (getline(file, line)) {
        string ten, tuoi, diachi, sdt;
        stringstream ss(line);
        ss >> ten >> tuoi >> diachi >> sdt;
        sinhVien sinhVien(ten, tuoi, diachi, sdt);
        dsThongTin.push_back(sinhVien);
    }
    file.close();
}

sinhVien::sinhVien(string basicString, string basicString1, string basicString2, string basicString3) {

}


#endif //P1_SINHVIEN_H
