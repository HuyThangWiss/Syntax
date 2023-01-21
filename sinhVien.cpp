//
// Created by huythang on 1/21/23.
//

#include "sinhVien.h"
#include <iostream>
#include <fstream>

using namespace std;



int main(){
    sinhVien sv;
  //  sv.nhap();
  //  sv.hienThi();
    sv.themThongTin();
    cout<<"*********** Thong tin *****"<<endl;
    for(int i=0;i<sv.dsThongTin.size();i++){
        sv.dsThongTin[i].hienThi();
    }
    cout<<"*********** Thong tin tim kiem *****"<<endl;
    sv.timKiemTheoTen("t1");
    cout<<"Ghi du lieu vao file"<<endl;
    sv.ghiDuLieuVaoFile();
}
