%undefine _missing_build_ids_terminate_build

%global repo            pmm-dump
%global provider        github.com/datacosmos-br/%{repo}
%global commit          dcc5549ac5ce927e2480daaad0e9e6e901ad6d52
%global shortcommit     %(c=%{commit}; echo ${c:0:7})
%define build_timestamp %(date -u +"%y%m%d%H%M")
%define release         1
%define rpm_release     %{release}.%{build_timestamp}.%{shortcommit}%{?dist}

Name:		pmm-dump
Version:	0.8.1
Release:	%{rpm_release}
Summary:	Percona PMM Dump allows to export and import monitoring metrics and query analytics.

License:	AGPLv3
URL:		https://%{provider}
Source0:	https://%{provider}/archive/%{commit}.tar.gz

%description
%{summary}

%prep
%setup -q -n %{repo}-%{commit}

%build
make build BRANCH="main" COMMIT="%{shortcommit}" VERSION="%{version}"

%install
install -d -p %{buildroot}%{_sbindir}
install -p -m 0755 pmm-dump %{buildroot}%{_sbindir}/pmm-dump

%files
%license LICENSE
%doc README.md
%{_sbindir}/pmm-dump


%changelog
* Mon Jun 15 2026 Marlon Costa <marlon.costa@datacosmos.com.br> - 0.8.1-dc2
- Pin pmm-dump to datacosmos-br fork tag v0.8.1-dc2 (commit dcc5549) as SSOT
- Aligns the committed spec with pmm-submodules/sources/pmm-dump so offline
  rpmspec and the build-time fork-provider override converge

* Wed Apr 22 2026 Maxim Kondratenko <maxim.kondratenko@percona.com> - 3.7.1
- PMM-14441 Upgrade pmm-dump to support encryption

* Wed Sep 24 2025 Michael Okoko <michael.okoko@percona.com> - 3.4.1
- PMM-14349 Update pmm-dump sources.

* Mon Jul 28 2025 Michael Okoko <michael.okoko@percona.com> - 3.4.0
- PMM-14215 Default to main branch for pmm-dump
- PMM-14085 Fix an issue where pmm-dump would not export OS metrics when used via the PMM GUI

* Thu Aug 8 2024 Alex Demidoff <alexander.demidoff@percona.com> - 3.0.0-1
- PMM-13282 Migrate pmm-dump to v3 API

* Thu Nov 23 2023 Artem Gavrilov <artem.gavrilov@percona.com> - 0.7.0-ga
- PMM-12460 Update pmm-dump to v0.7.0-ga version

* Tue Mar 29 2022 Alex Tymchuk <alexander.tymchuk@percona.com> - 0.6.0-1
- Initial pmm-dump version
